// src/utils/navigation.ts
let navigationCallback: ((path: string) => void) | null = null;

export const registerNavigationCallback = (callback: (path: string) => void) => {
  navigationCallback = callback;
};

export const navigate = (path: string) => {
  if (navigationCallback) {
    navigationCallback(path);
  } else {
    console.warn('Navigation callback not registered');
    // Fallback to window.location if no callback is registered
    window.location.href = path;
  }
};