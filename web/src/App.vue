<template>
	<v-app>
		<Header v-if="isLogged" />
		<v-content>
			<v-container>
				<router-view></router-view>
			</v-container>
		</v-content>
	</v-app>
</template>

<script>
import { mapState } from 'vuex'
import Header from './components/Header.vue'

export default {
	components: {
		Header
	},
	data() {
		return {
			user: {}
		}
	},
	computed: mapState(['isLogged']),
	methods: {

	},
	mounted() {
		let userData = localStorage.getItem('__user__')
		if (!userData) {
			this.$router.push('/login')
		} else {
			// has user
			this.isLogged = true
		}

	},
}
</script>

<style></style>

