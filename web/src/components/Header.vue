<template>
    <v-toolbar app>
        <v-toolbar-title class="headline text-upercase mr-4">
            <span>Stock</span>
            <span class="font-weight-light">Trader</span>
        </v-toolbar-title>
        <v-toolbar-items>
            <v-btn flat to="/">Inicio</v-btn>
            <v-btn flat to="/stocks">Ações</v-btn>
            <v-btn flat to="/portfolio">Portfolio</v-btn>
        </v-toolbar-items>

        <v-spacer></v-spacer>

        <v-toolbar-items>
            <v-btn flat>Finalizar Dia</v-btn>
            <v-menu offset-y>
                <v-btn flat slot="activator">Salvar & Carregar</v-btn>
                <v-list>
                    <v-list-tile>
                        <v-list-tile-title>Salvar Dados</v-list-tile-title>
                    </v-list-tile>
                    <v-list-tile>
                        <v-list-tile-title>Carregar Dados</v-list-tile-title>
                    </v-list-tile>
                </v-list>
            </v-menu>
            <v-layout align-center>
                <span class="text-uppercase grey--text text-darken-2">
                    Saldo: {{ founds  }}
                </span>
            </v-layout>
        </v-toolbar-items>
    </v-toolbar>
</template>

<script>
import config from '@/config/config';

export default {
    data() {
        return {
            founds: 0
        }
    },
    methods: {
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