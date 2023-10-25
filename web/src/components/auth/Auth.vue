<template>
    <div class="auth-content">
        <div class="auth-modal">
            <div class="auth-title">{{ showSignup ? 'Cadastro' : 'Login' }}</div>

            <input type="text" v-if="showSignup" v-model="user.name" name="name" placeholder="Nome">
            <input type="email" v-model="user.email" name="email" placeholder="E-mail">
            <input type="password" v-model="user.password" name="password" placeholder="Senha">
            <input type="password" v-if="showSignup" v-model="user.confirmPassword" placeholder="Confirme a senha">
        
            <button v-if="showSignup" @click="signup">Registrar</button>
            <button v-else @click="signin">Entrar</button>

            <a href @click.prevent="showSignup = !showSignup">
                <span v-if="showSignup">Já tem cadastro? Acesse o Login!</span>
                <span v-else>Não tem cadastro? Registre-se aqui!</span>
            </a>
        </div>
    </div>
</template>

<script>
import config from '@/config/config'


export default {
    name: 'Auth',
    data: function() {
        return {
            showSignup: false,
            user: {}
        }
    },
    methods: {
        signin() {
            if(!this.user.email || !this.user.password) {
                alert('Crdenciais inválidas!')
                return
            }

            fetch(`${config.API_URL}/login`, {
                method: 'POST',
                body: JSON.stringify(this.user),
                headers: {
                    'Content-Type': 'application/json'
                }
            })
            .then(res => res.json())
            .then(res => {
                if(res.error) {
                    alert('Crdenciais inválidas!')
                    return
                }
                localStorage.setItem('__user__', JSON.stringify(res))
                this.$router.push('/')
            })
            .catch(err => {
                alert("ERROR::", err)
            })
        },
        signup() {
            // ...
        }
    },
}

</script>

<style>
    .auth-content {
        height: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
        margin-top: 50px;
    }

    .auth-modal {
        background-color: #fff;
        width: 350px;
        padding: 35px;
        box-shadow: 0 1px 5px rgba(0, 0, 0, 0.15);

        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .auth-title {
        font-size: 2.2em;
        font-weight: 100;
        margin-bottom: 15px;
    }

    .auth-modal input {
        border: 1px solid #bbb;
        width: 100%;
        margin-bottom: 14px;
        padding: 3px 8px;
        height: 40px;
        border-radius: 3px;
        outline: none;
    }

    .auth-modal button {
        align-self: flex-end;
        background-color: #2460ae;
        color: #fff;
        padding: 5px 15px;
        border: none;
        cursor: pointer;
        margin-bottom: 10px;
    }

</style>