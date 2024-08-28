// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  modules: ['@nuxtjs/tailwindcss'],
  routeRules: {
    "/api/v1/users/**": {
      proxy: "http://localhost:10002/api/v1/users/**"
    }
  }
})