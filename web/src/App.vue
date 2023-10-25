<template>
	<v-app>
		<Header v-if="showSignup" />
		<v-content>
			<v-container>
				<router-view></router-view>
			</v-container>
		</v-content>
	</v-app>
</template>

<script>
import Header from './components/Header.vue'

export default {
	components: {
		Header
	},
	data() {
		return {
			showSignup: false,
			user: {}
		}
	},
	methods: {

	},
	mounted() {
		let userData = localStorage.getItem('__user__')
		if (!userData) {
			this.$router.push('/login')
		} else {
			// has user
			this.showSignup = true
		}

	},
}
</script>

<style></style>

