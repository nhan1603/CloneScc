import usePremiseParam from './usePremiseParam'
import { useSearchParams } from 'react-router-dom'

jest.mock('react-router-dom', () => {
  return {
    ...jest.requireActual('react-router-dom'),
    useSearchParams: jest.fn()
  }
})

describe('usePremiseParam', () => {
  it('usePremiseParam can be show premiseId and handle setPremiseParam', () => {
    const searchParams = {
      get: () => 123
    }
    const setSearchParams = jest.fn()
    useSearchParams.mockReturnValue([
      searchParams,
      setSearchParams
    ])

    const { setPremiseParam, withPremiseParam } = usePremiseParam()
    setPremiseParam(123)
    const withPremiseParamWithID = withPremiseParam('/hello')
    expect(usePremiseParam().premiseId).toEqual(123)
    expect(withPremiseParamWithID).toEqual('/hello?premiseId=123')
  })

  it('usePremiseParam not have premiseId and handle withPremiseParam with url only', () => {
    const searchParams = {
      get: () => false
    }
    const setSearchParams = jest.fn()
    useSearchParams.mockReturnValue([
      searchParams,
      setSearchParams
    ])

    const { setPremiseParam, withPremiseParam } = usePremiseParam()
    setPremiseParam(false)
    const withPremiseParamWithID = withPremiseParam('/hello')
    expect(usePremiseParam().premiseId).toEqual(false)
    expect(withPremiseParamWithID).toEqual('/hello')
  })
})
