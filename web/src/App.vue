<template>
	<v-app>
		<Header />
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
	created() {
		this.$store.dispatch('initStocks')
		// this.$store.dispatch('loadPortfolio')
	},
}
</script>

<style>

</style>

