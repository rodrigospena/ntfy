// This is a separate file since the other utils import `config.js`, which depends on `window`
// and cannot be used in the service worker

import emojisMapped from "./emojisMapped";
import { ACTION_HTTP, ACTION_VIEW } from "./actions";

const toEmojis = (tags) => {
  if (!tags) return [];
  return tags.filter((tag) => tag in emojisMapped).map((tag) => emojisMapped[tag]);
};

export const formatTitle = (m) => {
  const emojiList = toEmojis(m.tags);
  if (emojiList.length > 0) {
    return `${emojiList.join(" ")} ${m.title}`;
  }
  return m.title;
};

const formatTitleWithDefault = (m, fallback) => {
  if (m.title) {
    return formatTitle(m);
  }
  return fallback;
};

export const formatMessage = (m) => {
  if (m.title) {
    return m.message || "";
  }
  const emojiList = toEmojis(m.tags);
  if (emojiList.length > 0) {
    return `${emojiList.join(" ")} ${m.message || ""}`;
  }
  return m.message || "";
};

export const imageRegex = /\.(png|jpe?g|gif|webp|svg)(\b|\?|#|\/)/i;

export const getAttachmentUrl = (attachment) => {
  if (!attachment) return undefined;
  if (typeof attachment === "string") return attachment;
  return attachment.url;
};

export const isImage = (attachment) => {
  if (!attachment) return false;
  if (typeof attachment === "string") {
    return !!attachment.match(imageRegex);
  }
  if (attachment.type) {
    return attachment.type.startsWith("image/");
  }
  return !!(attachment.name?.match(imageRegex) || attachment.url?.match(imageRegex));
};

export const icon = "/static/images/ntfy.png";
export const badge = "/static/images/ntfy-mask.svg";

/**
 * Computes a unique notification tag scoped by baseUrl, topic, and sequence ID.
 * This ensures notifications from different topics with the same sequence ID don't collide.
 */
export const notificationTag = (baseUrl, topic, sequenceId) => `${baseUrl}/${topic}/${sequenceId}`;

export const toNotificationParams = ({ message, defaultTitle, topicRoute, baseUrl, topic }) => {
  const attachment = message.attachment || message.attach;
  const attachmentUrl = getAttachmentUrl(attachment);
  const image = isImage(attachment) ? attachmentUrl : undefined;
  const sequenceId = message.sequence_id || message.id;
  const tag = notificationTag(baseUrl, topic, sequenceId);
  const subscriptionId = `${baseUrl}/${topic}`;

  const iconUrl = message.icon || (isImage(attachment) ? attachmentUrl : icon);

  // https://developer.mozilla.org/en-US/docs/Web/API/Notifications_API
  return [
    formatTitleWithDefault(message, defaultTitle),
    {
      body: formatMessage(message),
      badge,
      icon: iconUrl,
      image,
      timestamp: message.time * 1000,
      tag, // Scoped by baseUrl/topic/sequenceId to avoid cross-topic collisions
      renotify: true,
      silent: false,
      // This is used by the notification onclick event
      data: {
        subscriptionId,
        message,
        topicRoute,
        url: message.click,
      },
      actions: message.actions
        ?.filter(({ action }) => action === ACTION_VIEW || action === ACTION_HTTP)
        .map(({ label }) => ({
          action: label,
          title: label,
        })),
    },
  ];
};

export const messageWithSequenceId = (message) => {
  if (message.sequenceId) {
    return message;
  }
  return { ...message, sequenceId: message.sequence_id || message.id };
};
