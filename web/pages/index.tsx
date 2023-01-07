import useSWR from 'swr'
import Link from 'next/link'

const rightPointerIcon = (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    fill='none'
    viewBox='0 0 24 24'
    strokeWidth={1.5}
    stroke='currentColor'
    className='w-10 border rounded-lg shadow-md p-2 hover:cursor-pointer active:bg-gray-200 text-slate-700 hover:text-slate-900 active:text-slate-900'
  >
    <path
      strokeLinecap='round'
      strokeLinejoin='round'
      d='M13.5 4.5L21 12m0 0l-7.5 7.5M21 12H3'
    />
  </svg>
)

const fetcher = (url: string) => fetch(url).then((res) => res.json())

export default function Home() {
  const { data: services } = useSWR('/api/services', fetcher)

  return (
    <div className='flex flex-col'>
      <div className='overflow-x-auto sm:-mx-6 lg:-mx-8'>
        <div className='py-2 inline-block min-w-full sm:px-6 lg:px-8'>
          <div className='overflow-hidden'>
            <table className='min-w-full'>
              <thead className='border-b'>
                <tr>
                  <th className='text-sm font-medium text-gray-900 px-6 py-4 text-left'>
                    #
                  </th>
                  <th className='text-sm font-medium text-gray-900 px-6 py-4 text-left'>
                    Name
                  </th>
                  <th className='text-sm font-medium text-gray-900 px-6 py-4 text-left'>
                    URL
                  </th>
                  <th className='text-sm font-medium text-gray-900 px-6 py-4 text-left'>
                    Last Deployed
                  </th>
                  <th className='text-sm font-medium text-gray-900 px-6 py-4 text-left'>
                    Action
                  </th>
                </tr>
              </thead>
              <tbody>
                {services?.data?.items?.map((elem: any, idx: number) => (
                  <tr className='border-b' key={idx}>
                    <td className='px-6 py-2 whitespace-nowrap text-sm font-medium text-gray-900'>
                      {idx + 1}
                    </td>
                    <td className='text-sm text-gray-900 font-light px-6 py-2 whitespace-nowrap'>
                      {elem.metadata?.name}
                    </td>
                    <td className='text-sm text-gray-900 font-light px-6 py-2 whitespace-nowrap'>
                      <Link
                        href={elem.status?.url}
                        className='text-blue-700 ease-linear duration-200 hover:text-blue-900 hover:cursor-pointer'
                      >
                        {elem.status.url}
                      </Link>
                    </td>
                    <td className='text-sm text-gray-900 font-light px-6 py-2 whitespace-nowrap'>
                      {elem.spec?.template?.metadata?.annotations?.[
                        'client.knative.dev/updateTimestamp'
                      ] || elem.metadata?.creationTimestamp}
                    </td>
                    <td className='text-sm text-gray-900 font-light px-6 py-2 whitespace-nowrap'>
                      {rightPointerIcon}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  )
}
