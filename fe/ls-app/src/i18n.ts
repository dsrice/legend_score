import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import LanguageDetector from 'i18next-browser-languagedetector';

// Import translation files
import translationEN from './locales/en/translation.json';
import translationJA from './locales/ja/translation.json';

// Resources for translations
const resources = {
  en: {
    translation: translationEN
  },
  ja: {
    translation: translationJA
  }
};

i18n
  // Detect user language
  .use(LanguageDetector)
  // Pass the i18n instance to react-i18next
  .use(initReactI18next)
  // Initialize i18next
  .init({
    resources,
    fallbackLng: 'en', // Default language if detection fails or language not supported
    detection: {
      // Order and from where user language should be detected
      order: ['navigator', 'htmlTag', 'path', 'subdomain'],
      // Only detect languages that are in the 'resources' object
      checkWhitelist: true
    },
    interpolation: {
      escapeValue: false // React already safes from XSS
    }
  });

export default i18n;