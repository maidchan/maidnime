import { createRouter, createWebHistory } from 'vue-router'
import HomeView from './views/HomeView.vue'
import VideoPlayerView from './views/VideoPlayerView.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView
        },
        {
            path: '/watch',
            name: 'watch',
            component: VideoPlayerView
        },
        {
            path: '/browse',
            name: 'browse',
            component: () => import('./views/BrowseView.vue')
        },
        {
            path: '/anime/:id',
            name: 'anime-detail',
            component: () => import('./views/AnimeDetailView.vue')
        },
        {
            path: '/favorites',
            name: 'favorites',
            component: () => import('./views/FavoritesView.vue'),
            meta: { hideNavbar: true }
        },
        {
            path: '/settings',
            name: 'settings',
            component: () => import('./views/SettingsView.vue'),
            meta: { hideNavbar: true }
        }
    ],
    scrollBehavior(_to, _from, savedPosition) {
        if (savedPosition) {
            return savedPosition
        } else {
            return { top: 0 }
        }
    }
})

export default router
