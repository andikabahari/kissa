import { NextApiResponse, NextApiRequest } from 'next'
import { KISSA_SERVER_URL } from '../../../utils/constants'

const url = new URL(`${KISSA_SERVER_URL}/api/services`)

export default async function handler(
  nextReq: NextApiRequest,
  nextRes: NextApiResponse
) {
  let serviceName: string
  let res: Response
  let body: any
  switch (nextReq.method) {
    case 'GET':
      serviceName = nextReq.query.service_name as string
      if (serviceName) url.searchParams.append('service_name', serviceName)
      res = await fetch(url)
      body = await res.json()
      nextRes.status(body.code).json(body)
      break
    case 'POST':
      res = await fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(nextReq.body),
      })
      body = await res.json()
      nextRes.status(body.code).json(body)
      break
    default:
      nextRes.status(404).json({ message: 'not found' })
  }
}
