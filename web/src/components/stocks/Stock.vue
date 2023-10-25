<template>
    <v-flex class="pr-3 pb-3" xs12 md6 lg4>
        <v-card class="green darken-3 white--text">
            <v-card-title class="headline">
                <strong>{{ stock.name }} <small>(Preço: {{ stock.price }})</small></strong>
            </v-card-title>
        </v-card>
        <v-card>
            <v-container fill-height>
                <v-text-field :error="insuficientFounds || !Number.isInteger(quantity)" v-model.number="quantity" label="Quantidade" type="number"></v-text-field>
                <v-btn 
                    @click="buyStock" 
                    :disabled="insuficientFounds || quantity <= 0 || !Number.isInteger(quantity)" 
                    class="green darken-3 white--text"
                >
                    {{ insuficientFounds ? 'Insuficiente' : 'Comprar'  }}
                </v-btn>
            </v-container>
        </v-card>
    </v-flex>
</template>

<script>
import config from '../../config/config'

export default {
    props: ['stock'],
    data() {
        return {
            quantity: 0,
            founds: 0
        }
    },

    computed: {
        insuficientFounds() {
            return this.quantity * this.stock.price > this.founds
        }
    },
    methods: {
        buyStock() {
            const user_id = '6536bd96345f7de7d8b9604c'
            const order = {
                stock_id: this.stock._id,
                user_id: user_id,
                quantity: this.quantity
            }

            fetch(`${config.API_URL}/buy-portfolio`, {
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
                    alert('Compra realizada com sucesso!')
                }
            })
            .catch(error => {
                alert('ERROR ao realizar a compra', error)
            })
            .finally(() => {
                this.quantity = 0
            })

        },
        getUserData() {
            const user_id = '6536bd96345f7de7d8b9604c'
            fetch(`${config.API_URL}/users/${user_id}`)
                .then(response => response.json())
                .then(data => {
                    this.founds = data.founds
                })
        }
    },
    created() {
        this.getUserData()
    }
}
</script>

<style>

</style>