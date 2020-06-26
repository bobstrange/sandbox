import { useEffect } from 'react'
import Router from 'next/router'
import useRequest from '../../hooks/use-request'

export default () => {
  const { executeRequest } = useRequest({
    url: '/api/users/signout',
    method: 'post',
    body: {},
    onSuccess: () => Router.push('/'),
  })

  useEffect(() => {
    executeRequest()
  }, [])

  return <div>Signing you out...</div>
}
