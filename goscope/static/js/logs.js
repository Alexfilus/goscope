let logOffset = 0;

const logTableHeaders = `
<tr>
	<th class="custom-td">Message</th>
	<th class="custom-td">Time</th>
	<th class="custom-td"></th>
</tr>
`;

async function getLogs(offset) {
    try {
        const response = await axios.get('/goscope/log-records', {
            params: {
                "offset": offset,
            }
        });
        return response.data
    } catch (error) {
        console.error(error);
    }
}

function fillLogTable(logData) {
    let logTable = document.getElementById("log-table");
    logTable.innerHTML = logTableHeaders;
    logData.forEach(function (item) {
        let requestMoment = item.time;
        let elapsed = secondsToString(now - requestMoment)
        logTable.innerHTML += `
            <tr class="text-center">
			    <td class="monospaced p-3 custom-td">${item.error}</td>
			    <td class="p-3 custom-td">${elapsed}</td>
                <td class="p-3 custom-td">
                    <a class="cursor-pointer" href="/goscope/log-records/${item.uid}" target="_blank" rel="noopener noreferrer">
                        ${viewMoreImage}
                    </a>
                </td>
            </tr>`;
    });
}

function increaseLogOffset() {
    logOffset += goscopeEntriesPerPage;
}

function decreaseLogOffset() {
    if (logOffset !== 0) {
        logOffset -= goscopeEntriesPerPage;
    }
}

let prevLogPage = document.getElementById("logs-prev-page");
prevLogPage.onclick = async function () {
    decreaseLogOffset();
    const data = await getLogs(logOffset)
    if (data !== null && data.length > 0) {
        fillLogTable(data);
    } else {
        increaseLogOffset();
    }
}
let nextLogPage = document.getElementById("logs-next-page");
nextLogPage.onclick = async function () {
    increaseLogOffset();
    const data = await getLogs(logOffset);
    if (data !== null && data.length > 0) {
        fillLogTable(data);
    } else {
        decreaseLogOffset();
    }
}

document.addEventListener("DOMContentLoaded", async function () {
    let logData = await getLogs(logOffset);
    fillLogTable(logData);
});