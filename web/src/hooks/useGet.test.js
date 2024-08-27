import useGet, { buildUrl } from './useGet'
import { renderHook } from '@testing-library/react-hooks'

jest.mock('data/api', () => {
  return {
    ...jest.requireActual('data/api'),
    get: jest.fn()
  }
})

describe('useGet', () => {
  it('useGet can get success response data', async () => {
    const { result } = renderHook(() => useGet('/hello', {id: 123}))
    expect(result.current.isLoading).toEqual(false)
  })

  it('buildUrl can get url and paramStringArr', async () => {
    const { result } = renderHook(() => buildUrl('/hello', {id: '123'}))
    expect(result.current).toEqual('/hello?&id=123')
  })

  it('buildUrl can get url only', async () => {
    const { result } = renderHook(() => buildUrl('/url-only'))
    expect(result.current).toEqual('/url-only')
  })
})
