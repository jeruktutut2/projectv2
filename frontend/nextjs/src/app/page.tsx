import Image from "next/image";
import Link from 'next/link';

export default function Home() {
    return (
        <div>
            <header className="fixed w-full top-0 left-0 px-2.5">
                <nav className="flex items-center justify-between mx-auto max-w-full py-[22px] px-0">
                    <a href="#" className="hover:text-gray-500">
                        <span>E-COMMERSE</span>
                    </a>
                    <ul className="flex list-none gap-[35px]">
                        <li><a href="menu1" className="no-underline text-[1.1rem] font-medium hover:text-gray-500">MENU 1</a></li>
                        <li><a href="menu2" className="no-underline text-[1.1rem] font-medium hover:text-gray-500">MENU 2</a></li>
                        <li><a href="menu3" className="no-underline text-[1.1rem] font-medium hover:text-gray-500">MENU 3</a></li>
                    </ul>
                    <Link href={`/login`}><button className="hover:text-black border border-gray px-4 py-2 rounded hover:bg-gray-200 active:bg-gray-300 active:border-gray-500">LOGIN</button></Link>
                </nav>
            </header>
        </div>
    );
}
