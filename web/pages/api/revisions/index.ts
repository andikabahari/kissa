import { NextApiResponse, NextApiRequest } from 'next'
import { KISSA_SERVER_URL } from '../../../utils/constants'

export default async function handler(
  nextReq: NextApiRequest,
  nextRes: NextApiResponse
) {
  try {
    const url = new URL(`${KISSA_SERVER_URL}/api/revisions`)
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
      default:
        nextRes.status(404).json({ message: 'not found' })
    }
  } catch (e) {
    nextRes.status(500).json({ message: 'internal server error' })
  }
}
