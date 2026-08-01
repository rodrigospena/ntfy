import * as React from "react";
import { createRoot } from "react-dom/client";
import App from "./components/App";
import registerSW from "./registerSW";

registerSW();

if ("Notification" in window && Notification.permission === "default") {
  window.addEventListener("load", () => {
    Notification.requestPermission();
  });
}

const root = createRoot(document.querySelector("#root"));
root.render(<App />);
