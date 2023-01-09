import useSWR from 'swr'
import Link from 'next/link'
import { Button, Table } from 'flowbite-react'
import Head from 'next/head'

const fetcher = (url: string) => fetch(url).then((res) => res.json())

export default function Home() {
  const { data: services } = useSWR('/api/services', fetcher)

  return (
    <>
      <Head>
        <title>Kissa</title>
      </Head>
      <div className='overflow-y-auto max-h-[576px] mt-8'>
        <Table hoverable={true}>
          <Table.Head>
            <Table.HeadCell>Service</Table.HeadCell>
            <Table.HeadCell>Last Deployed</Table.HeadCell>
            <Table.HeadCell>
              <span className='sr-only'>Details</span>
            </Table.HeadCell>
          </Table.Head>
          <Table.Body className='divide-y'>
            {services?.data?.items?.map((val: any, idx: number) => (
              <Table.Row
                key={idx}
                className='bg-white dark:border-gray-700 dark:bg-gray-800'
              >
                <Table.Cell className='whitespace-nowrap font-medium text-gray-900 dark:text-white'>
                  <Link
                    href={val.status?.url}
                    className='text-blue-500 ease-linear duration-200 hover:text-blue-700 hover:cursor-pointer'
                  >
                    {val.metadata?.name}
                  </Link>
                </Table.Cell>
                <Table.Cell>
                  {val.spec?.template?.metadata?.annotations?.[
                    'client.knative.dev/updateTimestamp'
                  ] || val.metadata?.creationTimestamp}
                </Table.Cell>
                <Table.Cell>
                  <Link
                    href={`/services/details/${val.metadata?.name}`}
                    className='text-blue-500 ease-linear duration-200 hover:text-blue-700 hover:cursor-pointer'
                  >
                    Details
                  </Link>
                </Table.Cell>
              </Table.Row>
            ))}
          </Table.Body>
        </Table>
      </div>
    </>
  )
}
