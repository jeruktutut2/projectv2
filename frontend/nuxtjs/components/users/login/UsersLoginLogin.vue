<script setup>
    import { reactive, ref } from 'vue';
    import { getHttpClient } from "../../../commons/setups/axios";

    const requestBody = reactive({
        email: "",
        password: ""
    })
    const emailErrorMessage = ref("")
    const passwordErrorMessage = ref("")
    const messageErrorMessage = ref("")

    async function login() {
        try {
            const httpClient = getHttpClient()

            const response = await httpClient.post("/api/v1/users/login", requestBody, {
                headers: {
                    "Accept": "application/json",
                    "X-REQUEST-ID": "requestId"
                }
            })
            await navigateTo('/')
        } catch(error) {
            error.response.data.errors.forEach(element => {
                if (element.field === "email") {
                    emailErrorMessage.value = element.message
                } else if (element.field === "password") {
                    passwordErrorMessage.value = element.message
                } else if (element.field === "message") {
                    messageErrorMessage.value = element.message
                }
            });
        }
    }
    
</script>

<template>
    <div class="fixed top-1/2 left-1/2 max-w-[720px] w-full transform -translate-x-1/2 -translate-y-1/2 border-2">
        <div class="flex flex-wrap md:flex-nowrap flex-col md:flex-row w-full">
            <div class="flex text-center flex-col items-center justify-center w-full p-[35px]
                        border">
                <h2 class="text-center mb-[20px]">Welcome Back</h2>
                <p>Please login using your personal information to stay connected with us.</p>
            </div>
            <div class="w-full p-[35px]
                        border">
                <h2 class="text-center mb-[10px]">LOGIN</h2>
                <div>
                    <input type="text" id="email" name="email" placeholder="Email"
                            v-model="requestBody.email" 
                            class="h-[40px] w-full outline-none text-[0.80rem] text-stone-950 px-4 rounded-[3px] border-2 border-[#717171]"/>
                    <p class="text-red-300 text-[0.60rem]" v-if="emailErrorMessage !== ''">{{ emailErrorMessage }}</p>
                </div>
                <div>
                    <input type="password" id="password" name="password" placeholder="Password" 
                            v-model="requestBody.password"
                            class="h-[40px] w-full outline-none text-[0.80rem] text-stone-950 px-4 rounded-[3px] mt-2 border-2 border-[#717171]"/>
                            <p class="text-red-300 text-[0.60rem]" v-if="passwordErrorMessage !== ''">{{ passwordErrorMessage }}</p>
                </div>
                <a href="#" className="text-gray-500 no-underline hover:underline inline-flex mt-2 text-[0.70rem]">Forgot Password</a>
                <button type="button"
                        class="h-[40px] w-full outline-none border-0 text-[0.80rem] font-semibold rounded-[3px] mt-1 text-white cursor-pointer bg-blue-500 hover:bg-blue-400 active:bg-blue-600 disabled:bg-blue-700"
                        @click="login">Login</button>
                <div className="text-center text-[0.70rem] mt-2">
                    Don&apos;t have an account? &nbsp;
                    <a href="#" className="text-gray-500 no-underline hover:underline">Sign Up</a>
                </div>
            </div>
        </div>
    </div>
</template>