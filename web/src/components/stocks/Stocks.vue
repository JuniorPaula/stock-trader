<template>
    <v-layout row wrap>
        <Stock v-for="stock in stocks" :key="stock.id" :stock="stock" />
    </v-layout>
</template>

<script>
import config from '../../config/config';
import Stock from './Stock.vue';

export default {
    components: { Stock },
    data() {
        return {
            stocks: []
        }
    },
    methods: {
        getStocks() {
            const useData = localStorage.getItem('__user__')
            const user = JSON.parse(useData)

            fetch(`${config.API_URL}/api/stocks`, {
                headers: {
                    'Authorization': `Bearer ${user.token}`
                }
            })
                .then(response => response.json())
                .then(data => {
                    this.stocks = data
                })
        }
    },
    created() {
        this.getStocks()
    }
}
</script>

<style>

</style>