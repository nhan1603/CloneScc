import Button from 'components/button'
import FormFieldInput from 'components/form/form-field-input'
import { useState } from 'react'
import { FormProvider, useForm } from 'react-hook-form'
import { useNavigate } from 'react-router'
import './LoginForm.scss'
import { setToken } from 'data/auth'
import { zodResolver } from '@hookform/resolvers/zod'
import loginSchema from 'schemas/loginSchema'
import api from 'data/api'

const LoginForm = () => {
  const [error, setError] = useState()
  const navigate = useNavigate()
  const [isLoading, setIsLoading] = useState(false)
  const form = useForm({
    mode: 'all',
    defaultValues: {
      username: '',
      password: ''
    },
    resolver: zodResolver(loginSchema)
  })

  const onSubmit = async ({ username, password }) => {
    try {
      setIsLoading(true)
      setError(null)
      const res = await api().post('/api/public/v1/auth/ou', {
        email: username,
        password
      })

      if (res.data?.token) {
        setToken(res.data.token)
        setTimeout(() => {
          navigate('/')
        }, 0)
      } else {
        setError('Something went wrong')
      }
    } catch (err) {
      setError('Invalid username or password')
      console.error(err)
    } finally {
      setIsLoading(false)
    }
  }
  return (
    <FormProvider {...form}>
      <div>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <h1>Welcome to Security Command Center</h1>
          <div>
            Enhance control of security operations: monitor incident reports, track surveillance
            feeds, access alarm systems, and more...
          </div>
          {error && <div className="error">{error}</div>}
          <FormFieldInput required name="username" data-testid="username" placeholder="Username" />
          <FormFieldInput
            required
            name="password"
            data-testid="password"
            placeholder="Password"
            type="password"
          />
          <Button
            onClick={form.handleSubmit(onSubmit)}
            disabled={isLoading}
            buttonType="solid"
            data-testid="submit-login">
            Login
          </Button>
          {/* For using Enter button */}
          <input type="submit" hidden></input>
        </form>
      </div>
    </FormProvider>
  )
}

export default LoginForm
