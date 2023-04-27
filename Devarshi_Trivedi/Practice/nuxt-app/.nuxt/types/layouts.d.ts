import { ComputedRef, Ref } from 'vue'
export type LayoutKey = "default" | "products"
declare module "/home/devarshi.trivedi/training/training/Devarshi_Trivedi/Practice/nuxt-app/node_modules/nuxt/dist/pages/runtime/composables" {
  interface PageMeta {
    layout?: false | LayoutKey | Ref<LayoutKey> | ComputedRef<LayoutKey>
  }
}