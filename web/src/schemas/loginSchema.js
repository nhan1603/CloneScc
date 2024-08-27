import z from 'zod'

const { string, object } = z

const loginSchema = object({
  username: string().min('1', 'Username is a required field'),
  password: string().min('1', 'Password is a required field')
})

export default loginSchema
