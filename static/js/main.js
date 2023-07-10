function addCategory() {
    let name = document.getElementById("new-category-name").value;

    $.ajax({
        url: `http://www.catalog-pattern.com/c`,
        type: "POST",
        data: JSON.stringify({
            name: name
        })
    }).always(HandleOrToast);
}

function saveItem(){

    let name = document.getElementById("item-name").value;
    let price = parseFloat(document.getElementById("item-price").value);
    let description = document.getElementById("item-description").value;
    let image = document.getElementById("item-image").files[0];
    let category = parseInt(document.getElementById("item-category").value);

    var formData = new FormData();

    formData.append("name", name);
    formData.append("price", price);
    formData.append("description", description);
    formData.append("image", image);
    formData.append("category", category);

    $.ajax({
        url: `http://www.catalog-pattern.com/i`,
        type: "POST",
        data: formData,
        contentType: false,
        processData: false
    }).always(HandleOrToast);

}



//*****************************************************//
//*********************** Utils ***********************//
//*****************************************************//

function HandleOrToast(jqXHR, textStatus) {

    if (jqXHR.status === 300) {
        // header location contains the string URL to redirect to
        window.location.replace(jqXHR.getResponseHeader("LOCATION"));
        return
    } else {
        let message;
        try {
            message = JSON.parse(jqXHR.responseText).message;
        } catch (err) {
            message = `[${jqXHR.status}] ${jqXHR.responseText}`;
        }
        if (textStatus != "success") {
            console.log(`${textStatus}: ${message}`, 10 * 1000); // 4000 is the duration of the toast
        } else {
            location.reload();
        }
    }
}

function Reload(jqXHR, textStatus) {
    location.reload()
}