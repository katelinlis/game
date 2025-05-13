// https://nuxt.com/docs/api/configuration/nuxt-config
import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  ssr: false,
  vite: {
    plugins: [
      tailwindcss(),
    ],
  },

  app:{
    head:{
      script: [
        {src: "https://telegram.org/js/telegram-web-app.js?57"},
        {src: "https://s3.eu-central-1.amazonaws.com/cdn.telemetree.io/telemetree-pixel.js"}

      ]
    }
  },
  css: ['~/assets/css/main.css'],
  modules: [
    // ...
    '@pinia/nuxt',
  ],
})
