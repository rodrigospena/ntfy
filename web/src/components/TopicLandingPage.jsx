import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import {
  Box,
  Card,
  CardContent,
  Typography,
  Button,
  Avatar,
  Container,
  Alert,
  CircularProgress,
  Paper,
  Chip,
} from "@mui/material";
import NotificationsActiveIcon from "@mui/icons-material/NotificationsActive";
import NotificationsOffIcon from "@mui/icons-material/NotificationsOff";
import CheckCircleIcon from "@mui/icons-material/CheckCircle";
import NotificationsIcon from "@mui/icons-material/Notifications";
import config from "../app/config";
import subscriptionManager from "../app/SubscriptionManager";
import { subscribeTopic } from "./SubscribeDialog";
import hideSplash from "../app/splash";

const TopicLandingPage = () => {
  const { topic } = useParams();
  const [loading, setLoading] = useState(false);
  const [subscribed, setSubscribed] = useState(false);
  const [permission, setPermission] = useState(
    typeof window !== "undefined" && "Notification" in window ? Notification.permission : "default"
  );
  const [error, setError] = useState(null);

  useEffect(() => {
    hideSplash();
    checkSubscriptionStatus();
  }, [topic]);

  const checkSubscriptionStatus = async () => {
    try {
      const subs = await subscriptionManager.all();
      const isSub = subs?.some(
        (s) => s.topic === topic && s.baseUrl === config.base_url
      );
      setSubscribed(!!isSub);
    } catch (e) {
      console.error("Error checking subscription status", e);
    }
  };

  const handleSubscribe = async () => {
    setLoading(true);
    setError(null);
    try {
      if ("Notification" in window && Notification.permission !== "granted") {
        const perm = await Notification.requestPermission();
        setPermission(perm);
        if (perm === "denied") {
          setError("As notificações foram bloqueadas no navegador. Para receber avisos, altere as permissões do site.");
          setLoading(false);
          return;
        }
      }

      await subscribeTopic(config.base_url, topic, {});
      setSubscribed(true);
      if ("Notification" in window) {
        setPermission(Notification.permission);
      }
    } catch (e) {
      console.error("Error subscribing to topic", e);
      setError("Não foi possível realizar a inscrição. Tente novamente.");
    } finally {
      setLoading(false);
    }
  };

  const handleUnsubscribe = async () => {
    setLoading(true);
    setError(null);
    try {
      await subscriptionManager.remove(config.base_url, topic);
      setSubscribed(false);
    } catch (e) {
      console.error("Error unsubscribing", e);
      setError("Não foi possível cancelar a inscrição.");
    } finally {
      setLoading(false);
    }
  };

  // Capitalize and format topic name (e.g. cafecomfamilia -> Café Com Família)
  const formattedTopicName = topic
    ? topic
        .replace(/[-_]/g, " ")
        .replace(/\b\w/g, (char) => char.toUpperCase())
    : "Tópico";

  return (
    <Box
      sx={{
        minHeight: "100vh",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        background: (theme) =>
          theme.palette.mode === "light"
            ? "linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%)"
            : "linear-gradient(135deg, #1e1e2f 0%, #0f0f17 100%)",
        p: 2,
      }}
    >
      <Container maxWidth="xs">
        <Card
          elevation={8}
          sx={{
            borderRadius: 4,
            textAlign: "center",
            p: 2,
            backdropFilter: "blur(10px)",
            boxShadow: "0 8px 32px 0 rgba(0, 0, 0, 0.2)",
          }}
        >
          <CardContent sx={{ display: "flex", flexDirection: "column", alignItems: "center", gap: 2 }}>
            <Avatar
              sx={{
                width: 80,
                height: 80,
                bgcolor: subscribed ? "success.main" : "primary.main",
                boxShadow: 3,
              }}
            >
              {subscribed ? (
                <NotificationsActiveIcon sx={{ fontSize: 44 }} />
              ) : (
                <NotificationsIcon sx={{ fontSize: 44 }} />
              )}
            </Avatar>

            <Box>
              <Typography variant="h5" component="h1" fontWeight="bold" gutterBottom>
                {formattedTopicName}
              </Typography>
              <Chip
                label={`Tópico: ${topic}`}
                size="small"
                variant="outlined"
                color="secondary"
                sx={{ fontFamily: "monospace" }}
              />
            </Box>

            <Typography variant="body2" color="text.secondary">
              {subscribed
                ? "Você está inscrito! As notificações deste tópico serão enviadas diretamente para este dispositivo."
                : "Receba notificações instantâneas e avisos deste tópico diretamente no seu navegador."}
            </Typography>

            {permission === "denied" && (
              <Alert severity="warning" sx={{ width: "100%", textAlign: "left", mt: 1 }}>
                As notificações estão bloqueadas no seu navegador. Ative as permissões nas configurações do site.
              </Alert>
            )}

            {error && (
              <Alert severity="error" sx={{ width: "100%", textAlign: "left", mt: 1 }}>
                {error}
              </Alert>
            )}

            <Box sx={{ width: "100%", mt: 2 }}>
              {subscribed ? (
                <Box sx={{ display: "flex", flexDirection: "column", gap: 1.5 }}>
                  <Paper
                    variant="outlined"
                    sx={{
                      p: 1.5,
                      display: "flex",
                      alignItems: "center",
                      justifyContent: "center",
                      gap: 1,
                      borderColor: "success.main",
                      bgcolor: (theme) => (theme.palette.mode === "light" ? "#e8f5e9" : "#1b2e1e"),
                    }}
                  >
                    <CheckCircleIcon color="success" />
                    <Typography variant="body2" fontWeight="medium" color="success.main">
                      Inscrição Ativa
                    </Typography>
                  </Paper>
                  <Button
                    variant="text"
                    color="error"
                    size="small"
                    onClick={handleUnsubscribe}
                    disabled={loading}
                    startIcon={<NotificationsOffIcon />}
                  >
                    Cancelar inscrição
                  </Button>
                </Box>
              ) : (
                <Button
                  variant="contained"
                  color="primary"
                  size="large"
                  fullWidth
                  onClick={handleSubscribe}
                  disabled={loading}
                  startIcon={
                    loading ? <CircularProgress size={20} color="inherit" /> : <NotificationsActiveIcon />
                  }
                  sx={{
                    py: 1.5,
                    borderRadius: 3,
                    fontSize: "1rem",
                    fontWeight: "bold",
                    textTransform: "none",
                    boxShadow: 3,
                    "&:hover": {
                      boxShadow: 6,
                    },
                  }}
                >
                  {loading ? "Inscrevendo..." : "Quero Receber Notificações"}
                </Button>
              )}
            </Box>
          </CardContent>
        </Card>
      </Container>
    </Box>
  );
};

export default TopicLandingPage;
