export default async function updateTargetQty(url:string, startDate:string,endDate:string, lines:string) {
    const headers = new Headers();
    headers.append(`Content-Type`, `application/json`);
    headers.append(`Accept`, `application/json`);

    const body = JSON.stringify({
        start_date: startDate,
        end_date: endDate,
        lines: lines
    });

    const options = {
        method: `POST`,
        headers: headers,
        body: body,
    }

    const request = await fetch(`${url}update_target_quantity`, options);
    const response = await request.json();
    if(import.meta.env.DEV) console.log(`update target qty response`, response);
    return response;
}