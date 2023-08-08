import useSWR from 'swr'

export type ApiResponse = {
  message: string
  data?: any
}

const baseUrl = process.env.NEXT_PUBLIC_SERVER_URL

const fetcher = async (url: string) => {
  const res = await fetch(url)
  const body = (await res.json()) as ApiResponse
  if (!res.ok) throw Error(body.message)
  return body.data
}

export const getServices = () => {
  const url = new URL('api/services', baseUrl)
  return useSWR(url.toString(), fetcher)
}

export const getService = (serviceName: string) => {
  const url = new URL('api/services/' + serviceName, baseUrl)
  return useSWR(url.toString(), fetcher)
}

export const createService = async (body: any) => {
  const url = new URL('api/services', baseUrl)
  const res = await fetch(url.toString(), {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body),
  })
  return res
}

export const getRevisions = (serviceName?: string) => {
  const url = new URL('api/revisions', baseUrl)
  if (serviceName) url.searchParams.set('service_name', serviceName)
  return useSWR(url.toString(), fetcher)
}
