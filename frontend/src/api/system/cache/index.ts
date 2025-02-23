import request from '/frontend/src/utils/request';

// 清除缓存
export function removeCache() {
    return request({
        url: '/api/v1/system/cache/remove',
        method: 'delete'
    })
}
