import Home from './components/Home.vue'
import Users from './components/Users.vue'
import FileSystem from './components/FileSystem.vue'
import System from './components/System.vue'
import Backup from './components/Backup.vue'
import Devices from './components/Devices.vue'
import Terminal from './components/Terminal.vue'
import Help from './components/Help.vue'

export default [
    {path: '/', component: Home},
    {path: '/users', component: Users},
    {path: '/filesystem', component: FileSystem},
    {path: '/system', component: System},
    {path: '/backup', component: Backup},
    {path: '/devices', component: Devices},
    {path: '/terminal', component: Terminal},
    {path: '/help', component: Help}
]