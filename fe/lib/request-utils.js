
export function headers(token){
    if(token)
        return new Headers({
        'Authorization': 'Bearer ' + token, 
        'Content-Type': 'application/json'
        }); 

    return new Headers({
        'Content-Type': 'application/json'
    });
};

export const get = async props => {
try {
    const res = await fetch(props.endpoint, {
    method: 'GET',
    headers: headers(props.token), 
    });
    check_token(res.status);
    if(res.status != 200) return [];
    return await res.json(); 
} catch { return []; }
};

export const post = async props => {

try {
    const res = await fetch(props.endpoint, {
    method: 'POST',
    headers: headers(props.token),
    body : JSON.stringify(props.data)  
    });

    check_token(res.status);

    return {
    res : await res.json(),
    status : res.status
    };

} catch { 
    return {
    res : null,
    status : -1
    }; 
}

};

export const del = async props => {

try {
    const res = await fetch(props.endpoint, {
    method: 'DELETE',
    headers: headers(props.token), 
    });
    
    check_token(res.status);

    if(res.status != 204) return false;
    
    return true;

} catch (error) {
    return false;
}

};

export const mod = async props => {
try {
    const res = await fetch(props.endpoint, {
    method: 'PUT',
    headers: headers(props.token), 
    body : JSON.stringify(props.data)  
    });

    check_token(res.status);
    
    if (res.status == 200)
    return {
        res : null,
        status : 200
    };

    return {
    res : await res.json(),
    status : res.status
    };
    
} catch (error) {
    return {
    res : null,
    status : -1
    };
}
};