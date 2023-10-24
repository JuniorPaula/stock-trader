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
            fetch(`${config.API_URL}/stocks`)
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