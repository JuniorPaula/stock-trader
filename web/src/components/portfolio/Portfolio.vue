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
        getUserStocks() {
            const user_id = '6536bd96345f7de7d8b9604c'
            fetch(`${config.API_URL}/users/${user_id}`)
                .then(response => response.json())
                .then(data => {
                    this.stocks = data.portfolios

                    fetch(`${config.API_URL}/stocks`)
                        .then(response => response.json())
                        .then(data => {
                            this.stocks = this.stocks.map(stock => {
                                const stockData = data.find(item => item._id === stock.stock_id)
                                return {
                                    ...stock,
                                    ...stockData
                                }
                            })
                        })
                })
        }
    },
    created() {
        this.getUserStocks()
    }

}
</script>

<style></style>