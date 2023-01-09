import { Button, Navbar as FlowbiteNavbar } from 'flowbite-react'
import Link from 'next/link'
import styles from './Navbar.module.css'

export default function Navbar() {
  return (
    <div className={styles.Navbar}>
      <div className='border-b mb-4'>
        <div className='container max-w-screen-lg mx-auto px-6'>
          <FlowbiteNavbar fluid={true} rounded={false} className='!px-0'>
            <Link href='/'>
              <span className='self-center whitespace-nowrap text-xl font-semibold text-blue-700 dark:text-white'>
                Kissa
              </span>
            </Link>
            <Link href='/services/deploy'>
              <Button gradientDuoTone='cyanToBlue'>Deploy 🚀</Button>
            </Link>
          </FlowbiteNavbar>
        </div>
      </div>
    </div>
  )
}
