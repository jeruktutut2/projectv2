"use client"

import { getHttpClient } from "@/comons/setups/axios"
import { AxiosError } from "axios"
import { useRouter } from "next/navigation"
import React, { useState } from "react"

export function LoginPage() {
    const router = useRouter()
    const [emailValue, setEmailOnChange] = useState("")
    const [passwordValue, setPasswordOnChange] = useState("")
    const [emailErrorMessage, setEmailErrorMessage] = useState("")
    const [passwordErrorMessage, setPasswordErrorMessage] = useState("")
    const [messageErrorMessage, setMessageErrorMessage] = useState("") 
    const [message, setMessage] = useState("")
    const [pending, setPending] = useState(false)

    const setUsername = (e: React.ChangeEvent<HTMLInputElement>) => {
        setEmailOnChange(e.target.value)
        setEmailErrorMessage("")
    }

    const setPassword = (e: React.ChangeEvent<HTMLInputElement>) => {
        setPasswordOnChange(e.target.value)
        setPasswordErrorMessage("")
    }

    const Login = async (event: React.MouseEvent<HTMLButtonElement>) => {
        // event: React.MouseEvent<HTMLButtonElement>
        // event.preventDefault();
        setPending(true)
        try {        
            const httpClient = getHttpClient()
            const requesstBody = {
                email: emailValue,
                password: passwordValue
            }
            const response = await httpClient.post("/api/v1/users/login", requesstBody, {
                headers: {
                    "Accept": "application/json",
                    "X-REQUEST-ID": "requestId"
                }
            })
            router.push("/")
        } catch(error: unknown) {
            if (error instanceof AxiosError) {
                error.response?.data.errors.forEach((element: any) => {
                    if (element.field === "email") {
                        setEmailErrorMessage(element.message)
                    } else if (element.field === "password") {
                        setPasswordErrorMessage(element.message)
                    } else if (element.field === "message") {
                        setMessageErrorMessage(element.message)
                    }
                });
            } else {
                console.log("error:", error);
            }
        }
        setPending(false)
    }

    return (
        <div className="fixed top-1/2 left-1/2 max-w-[720px] w-full transform -translate-x-1/2 -translate-y-1/2 border-2">
            <div className="flex flex-wrap md:flex-nowrap flex-col md:flex-row w-full">
                <div className="flex text-center flex-col items-center justify-center w-full p-[35px]
                                border">
                    <h2 className="text-center mb-[20px]">Welcome Back</h2>
                    <p>Please login using your personal information to stay connected with us.</p>
                </div>
                <div className="w-full p-[35px]
                                border">
                    <h2 className="text-center mb-[10px]">LOGIN</h2>
                    {message && <p className="text-center text-green-800 text-[0.80rem] mb-[15px]">{message}</p>}
                    {messageErrorMessage && <p className="text-center text-red-800 text-[0.80rem] mb-[15px]">{messageErrorMessage}</p>}
                    <div>
                        <input type="text" id="email" name="email" placeholder="Email" 
                                disabled={pending} 
                                value={emailValue}
                                onChange={(e) => setUsername(e)}
                                className={"h-[40px] w-full outline-none text-[0.80rem] text-stone-950 px-4 rounded-[3px] border-2 " + (emailErrorMessage === "" ? " border-[#717171]" : " border-red-500 ") }/>
                        {emailErrorMessage && <p className="text-red-300 text-[0.60rem]">{emailErrorMessage}</p>}
                    </div>
                    <div>
                        <input type="password" id="password" name="password" placeholder="Password" 
                                disabled={pending} 
                                value={passwordValue}
                                onChange={(e) => setPassword(e)}
                                className={"h-[40px] w-full outline-none text-[0.80rem] text-stone-950 px-4 rounded-[3px] mt-2 border-2 " + (passwordErrorMessage === "" ? " border-[#717171]" : " border-red-500 ")}/>
                        {passwordErrorMessage && <p className="text-red-300 text-[0.60rem]">{passwordErrorMessage}</p>}
                    </div>
                    <a href="#" className="text-gray-500 no-underline hover:underline inline-flex mt-2 text-[0.70rem]">Forgot Password</a>
                    <button type="button" disabled={pending} 
                            onClick={Login}
                            className="h-[40px] w-full outline-none border-0 text-[0.80rem] font-semibold rounded-[3px] mt-1 text-white cursor-pointer bg-blue-500 hover:bg-blue-400 active:bg-blue-600 disabled:bg-blue-700">{pending ? "Logging in..." : "Login"}</button>
                    <div className="text-center text-[0.70rem] mt-2">
                        Don&apos;t have an account? &nbsp;
                        <a href="#" className="text-gray-500 no-underline hover:underline">Sign Up</a>
                    </div>
                </div>
            </div>
        </div>
    )
}