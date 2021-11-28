import React from 'react'

interface Props {}

const Login = (props: Props) => {
  return (
    <div className="container mx-auto flex justify-center items-center w-100 h-screen">
      <div className="flex flex-col max-w-md w-full space-y-8 h-max p-8 bg-gray-100 border border-dark-600 rounded">
        <h1 className="font-bold text-3xl ">Sign in to your account</h1>
        <form className="w-full space-y-6">
          <div id="textField" className="flex flex-col w-100">
            <label htmlFor="email" className="text-sm">
              Email
            </label>
            <input
              className=" h-10 p-4"
              name="email"
              id="email"
              placeholder="Masukkan email anda"
            />
          </div>
          <div id="textField" className="flex flex-col mt-4">
            <label htmlFor="password" className="text-sm">
              Password
            </label>
            <input
              className=" h-10 p-4"
              name="password"
              id="password"
              placeholder="Masukkan password anda"
            />
          </div>
        </form>
      </div>
    </div>
  )
}

export default Login
