<template>
    <v-layout row wrap>
        <!-- :updateStocks listen event from Stock child component-->
        <Stock :updateStocks="getUserStocks" v-for="stock in stocks" :key="stock.id" :stock="stock" />
    </v-layout>
</template>

<script>
import config from '../../config/config';
import Stock from './Stock.vue';

export default {
    components: { Stock },
    data() {
        return {
            stocks: [],
            user: {}
        }
    },
    methods: {
        getUserStocks() {
            fetch(`${config.API_URL}/api/users/${this.user.user_id}`, {
                headers: {
                    'Authorization': `Bearer ${this.user.token}`
                }
            })
            .then(response => response.json())
            .then(data => {
                this.stocks = data.portfolios

                fetch(`${config.API_URL}/api/stocks`, {
                    headers: {
                        'Authorization': `Bearer ${this.user.token}`
                    }
                })
                .then(response => response.json())
                .then(data => {
                    this.stocks = this.stocks.map(stock => {
                        const stockData = data.find(item => item._id === stock.stock_id)
                        return {
                            portfolio_id: stock._id,
                            ...stock,
                            ...stockData
                        }
                    })
                })
            })
        }
    },
    created() {
        const userData = localStorage.getItem('__user__')
        this.user = JSON.parse(userData)

        this.getUserStocks() 
    },

}
</script>

<style></style>