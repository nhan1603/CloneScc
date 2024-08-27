import api from 'data/api'
import { useCallback, useState } from 'react'

const usePost = () => {
  const [isLoading, setIsLoading] = useState(false)

  const post = useCallback(async (url, payload, onSuccess, onError) => {
    try {
      setIsLoading(true)
      const response = await api().post(url, payload)
      onSuccess?.(response)
    } catch (error) {
      console.error(error)
      onError?.(error)
    } finally {
      setIsLoading(false)
    }
  }, [])

  return {
    post,
    isLoading
  }
}
export default usePost
