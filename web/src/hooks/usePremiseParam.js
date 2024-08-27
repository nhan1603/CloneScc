import { useSearchParams } from 'react-router-dom'

const usePremiseParam = () => {
  const [searchParams, setSearchParams] = useSearchParams()

  const setPremiseParam = (premiseId) => {
    setSearchParams({
      ...searchParams,
      premiseId
    })
  }

  const premiseId = searchParams.get('premiseId')

  return {
    premiseId,
    setPremiseParam,
    withPremiseParam: (url) => {
      if (!premiseId) {
        return url
      }
      return `${url}?premiseId=${premiseId}`
    }
  }
}

export default usePremiseParam
