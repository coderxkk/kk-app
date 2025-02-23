import request from '/frontend/src/utils/request';


export function getSysInfo() {
    return request({
        url: '/api/v1/system/monitor/server',
        method: 'get'
    })
}
