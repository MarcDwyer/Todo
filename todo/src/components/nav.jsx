import React from 'react'
import { NavItem, Navbar } from 'react-materialize'


const Nav = () => (
    <Navbar brand='Todo App' right className="amber darken-1">
        <NavItem onClick={() => console.log('test click')}>Your todos</NavItem>
        <NavItem href='#'>Your account</NavItem>
</Navbar>
)

export default Nav;