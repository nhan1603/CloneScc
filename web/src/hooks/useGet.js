import api from 'data/api'
import { useEffect, useState } from 'react'

const useGet = (url, param) => {
  const [response, setReponse] = useState(null)
  const [error, setError] = useState(null)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    const fetchData = async () => {
      try {
        setIsLoading(true)
        const res = await api().get(buildUrl(url, param))
        setReponse(res)
      } catch (e) {
        setError(e)
      } finally {
        setIsLoading(false)
      }
    }

    fetchData()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [JSON.stringify(param), url])

  return { response, error, isLoading }
}

export const buildUrl = (url, param) => {
  if (!param) {
    return url
  }
  const keys = Object.keys(param)
  const paramStringArr = keys.map((k) => `${k}=${param[k]}`)
  return [`${url}?`, ...paramStringArr].join('&')
}

export default useGet
