import Dropdown from './user_components/user_dropdown'
import Link from 'next/link'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faGithub } from '@fortawesome/free-brands-svg-icons';
export default function Footer() {
  return (
    <div className="flex items-start justify-between flex-wrap bg-slate-600 text-slate-300 p-5 w-full">
        <div className='flex flex-row lg:flex-col w-full lg:w-1/3 justify-center items-center'>
          <ul>
            <li>
            <h1>Find us on GitHub:</h1>
              <a className='hover:text-white hover:cursor-pointer'>
                <FontAwesomeIcon icon={faGithub} className="pr-3"/>
                Edoardo Geraci
              </a>
            </li>
            <li>
              <a className='hover:text-white hover:cursor-pointer'>
                <FontAwesomeIcon icon={faGithub} className="pr-3"/>
                Alessandro Versari
              </a>
            </li>
            <li>
              <a className='hover:text-white hover:cursor-pointer'>
                <FontAwesomeIcon icon={faGithub} className="pr-3"/>
                Andrea Bedini
              </a>
            </li>
          </ul>
        </div>
    </div>
  )}