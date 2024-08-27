import axios from 'axios'
import { getToken } from './auth'
export const API_ENDPOINT = process.env.REACT_APP_API_ENDPOINT

const axiosInstance = axios.create({
  baseURL: API_ENDPOINT,
  headers: {
    'Content-Type': 'application/json'
  }
})

export default function () {
  const token = getToken()

  if (token) {
    axiosInstance.defaults.headers['Authorization'] = token
  } else {
    delete axiosInstance.defaults.headers['Authorization']
  }

  return axiosInstance
}
