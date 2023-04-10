// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    runtimeConfig: {
        // The private keys which are only available server-side
        apiSecret: '123',
        // Keys within public are also exposed client-side
        public: {
          apiBase: '/api'
        }
      },
      vite: {
        css: {
          preprocessorOptions: {
            scss: {
              additionalData: '@use "@/assets/_colors.scss" as *;'
            },
            sass: {
              additionalData: '@use "@/assets/_colors.sass" as *\n'
            }
          }
        }
      },
      app: {
        head: {
          charset: 'utf-8',
          viewport: 'width=device-width, initial-scale=1',
        }
      },
      css : [
        "~/node_modules/bootstrap/dist/css/bootstrap.min.css"
      ]
})
