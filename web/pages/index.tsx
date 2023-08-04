import Link from 'next/link'
import { Badge, Table } from 'flowbite-react'
import Head from 'next/head'
import { getServices } from '../utils/api'

export default function Home() {
  const { data: services, error, isLoading } = getServices()

  return (
    <>
      <Head>
        <title>Kissa</title>
      </Head>
      <div className='overflow-y-auto max-h-[576px] mt-8'>
        <Table hoverable={true}>
          <Table.Head>
            <Table.HeadCell>Service</Table.HeadCell>
            <Table.HeadCell>Ready</Table.HeadCell>
            <Table.HeadCell>Last Deployed</Table.HeadCell>
            <Table.HeadCell>
              <span className='sr-only'>Details</span>
            </Table.HeadCell>
          </Table.Head>
          <Table.Body className='divide-y'>
            {isLoading ? (
              <Table.Row>
                <Table.Cell colSpan={4} className='bg-white'>
                  feching data...
                </Table.Cell>
              </Table.Row>
            ) : error ? (
              <Table.Row>
                <Table.Cell colSpan={4} className='bg-white'>
                  {error.message}
                </Table.Cell>
              </Table.Row>
            ) : (
              services.map((val: any, idx: number) => (
                <Table.Row
                  key={idx}
                  className='bg-white dark:border-gray-700 dark:bg-gray-800'
                >
                  <Table.Cell className='whitespace-nowrap font-medium text-gray-900 dark:text-white'>
                    <Link
                      href={val.url}
                      className='text-blue-500 ease-linear duration-200 hover:text-blue-700 hover:cursor-pointer'
                    >
                      {val.name}
                    </Link>
                  </Table.Cell>
                  <Table.Cell>
                    <div className='inline-block'>
                      {val.ready ? (
                        <Badge color='success'>True</Badge>
                      ) : (
                        <Badge color='failure'>False</Badge>
                      )}
                    </div>
                  </Table.Cell>
                  <Table.Cell>{val.last_deployed}</Table.Cell>
                  <Table.Cell>
                    <Link
                      href={`/services/details/${val.name}`}
                      className='text-blue-500 ease-linear duration-200 hover:text-blue-700 hover:cursor-pointer'
                    >
                      Details
                    </Link>
                  </Table.Cell>
                </Table.Row>
              ))
            )}
          </Table.Body>
        </Table>
      </div>
    </>
  )
}
