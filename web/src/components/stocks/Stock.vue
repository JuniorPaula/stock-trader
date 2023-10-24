<template>
    <v-flex class="pr-3 pb-3" xs12 md6 lg4>
        <v-card class="green darken-3 white--text">
            <v-card-title class="headline">
                <strong>{{ stock.name }} <small>(Preço: {{ stock.price }})</small></strong>
            </v-card-title>
        </v-card>
        <v-card>
            <v-container fill-height>
                <v-text-field v-model.number="quantity" label="Quantidade" type="number"></v-text-field>
                <v-btn 
                    @click="buyStock" 
                    :disabled="quantity <= 0 || !Number.isInteger(quantity)" 
                    class="green darken-3 white--text"
                >
                    Comprar
                </v-btn>
            </v-container>
        </v-card>
    </v-flex>
</template>

<script>
export default {
    props: ['stock'],
    data() {
        return {
            quantity: 0
        }
    },
    methods: {
        buyStock() {
            const order = {
                stock_id: this.stock.id,
                user_id: 1,
                quantity: this.quantity
            }
            this.quantity = 0
        }
    },
}
</script>

<style>

</style>