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
                    :disabled="quantity <= 0 || !Number.isInteger(quantity) || quantity > stock.quantity" 
                    class="blue darken-3 white--text"
                >
                    Vender
                </v-btn>
            </v-container>
        </v-card>
    </v-flex>
</template>

<script>
import config from '../../config/config'

export default {
    props: ['stock', 'updateStocks'],
    data() {
        return {
            quantity: 0
        }
    },
    methods: {
        sellStock() {
            const order = {
                _id: this.stock.portfolio_id,
                user_id: this.stock.user_id,
                quantity: this.quantity
            }
            
            fetch(`${config.API_URL}/sell-portfolio`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(order)
            })
            .then(response => response.json())
            .then(data => {
                if(data.error) {
                    alert(data.error)
                } else {
                    this.updateStocks() // emit event to parent component when sell stock
                    alert('Venda realizada com sucesso!')
                }
            })
            .catch(error => {
                alert('ERROR ao realizar a venda', error)
            })
            .finally(() => {
                this.quantity = 0
            })
        }
    },
}
</script>

<style>

</style>