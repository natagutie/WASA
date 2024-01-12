import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import login from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'

import AccountViewer from '../views/AccountViewer.vue'
import SearchView from '../views/SearchView.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: login},
		{path: '/session', component: HomeView},
		{path: '/search', component: SearchView},
		{path: '/home', component: HomeView},
		{path: '/user/:username/account', component: AccountViewer},
		{path: '/profile', component: ProfileView},
	]
})

export default router
