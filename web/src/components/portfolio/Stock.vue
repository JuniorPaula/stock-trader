<template>
    <v-flex class="pr-3 pb-3" xs12 md6 lg4>
        <v-card class="blue darken-3 white--text">
            <v-card-title class="headline">
                <strong>
                    {{ stock.name }} 
                    <small>(Preço: {{ stock.price }}) | Qtde: {{ stock.quantity }}</small>
                </strong>
            </v-card-title>
        </v-card>
        <v-card>
            <v-container fill-height>
                <v-text-field v-model.number="quantity" label="Quantidade" type="number"></v-text-field>
                <v-btn 
                    @click="sellStock" 
                    :disabled="quantity <= 0 || !Number.isInteger(quantity)" 
                    class="blue darken-3 white--text"
                >
                    Vender
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
        sellStock() {
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