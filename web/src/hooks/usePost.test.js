import usePost from './usePost'
import { act, renderHook } from '@testing-library/react-hooks'

jest.mock('data/api', () => {
  return {
    ...jest.requireActual('data/api'),
    get: jest.fn()
  }
})

describe('usePost', () => {
  it('usePost can get success response data', async () => {
    const { result } = renderHook(() => usePost())
    console.error = jest.fn()
    act(() => {
      /* fire events that update state */
      result.current.post()
    });
    expect(result.current.isLoading).toEqual(false)
  })
})
