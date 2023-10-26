import Vue from 'vue'
import './plugins/vuetify'
import App from './App.vue'
import router from './router'
import store from '@/config/store'

Vue.config.productionTip = false

Vue.filter('currency', value => {
	return new Intl.NumberFormat('pt-BR', {
		style: 'currency',
		currency: 'BRL',
		minimumFractionDigits: 2
	}).format(value)
})

new Vue({
	store,
	router,
	render: h => h(App),
}).$mount('#app')
