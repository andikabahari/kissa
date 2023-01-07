import Link from 'next/link'

export default function Navbar() {
  return (
    <nav className='mb-5  border-gray-300 border-b'>
      <ul className='flex flex-row justify-center'>
        <Link
          href='/'
          className='text-gray-700 font-medium ease-linear duration-200 hover:text-gray-900 hover:cursor-pointer'
        >
          <li className='p-3 mx-2'>Kissa</li>
        </Link>
      </ul>
    </nav>
  )
}
