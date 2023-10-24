<template>
    <v-layout row wrap>
        <Stock v-for="stock in stocks" :key="stock.id" />
    </v-layout>
</template>

<script>
import Stock from './Stock.vue';

export default {
    components: { Stock },
    data() {
        return {
            stocks: [
                { id: 1, name: 'Google', price: 110 },
                { id: 2, name: 'BMW', price: 320 },
                { id: 3, name: 'NuBank', price: 90 },
                { id: 1, name: 'Apple', price: 160 },
            ]
        }
    }
}
</script>

<style>

</style>