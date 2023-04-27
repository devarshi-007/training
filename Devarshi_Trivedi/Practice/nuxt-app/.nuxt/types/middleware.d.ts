import type { NavigationGuard } from 'vue-router'
export type MiddlewareKey = string
declare module "/home/devarshi.trivedi/training/training/Devarshi_Trivedi/Practice/nuxt-app/node_modules/nuxt/dist/pages/runtime/composables" {
  interface PageMeta {
    middleware?: MiddlewareKey | NavigationGuard | Array<MiddlewareKey | NavigationGuard>
  }
}