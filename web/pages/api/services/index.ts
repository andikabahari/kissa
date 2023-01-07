import { NextApiResponse, NextApiRequest } from 'next'
import { KISSA_SERVER_URL } from '../../../utils/constants'

const url = new URL(`${KISSA_SERVER_URL}/api/services`)

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  switch (req.method) {
    case 'GET':
      const serviceName = req.query.service_name as string
      if (serviceName) url.searchParams.append('service_name', serviceName)
      const fetchReq = await fetch(url)
      const body = await fetchReq.json()
      res.status(body.code).json(body)
      break
    default:
      res.status(404).json({ message: 'not found' })
  }
}
