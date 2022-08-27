// url: https://script.google.com/macros/s/.../exec?barcode=6287029390129
// webapp: AKfycbzw0TKWycxeB5sx1wIefAiEHeYQt2mVuM-NAZTccxedhyntdv8FvcUteOZ2k03wRHGE
// check: https://script.google.com/macros/s/AKfycbzw0TKWycxeB5sx1wIefAiEHeYQt2mVuM-NAZTccxedhyntdv8FvcUteOZ2k03wRHGE/exec?barcode=6287029390129
fileID = "1M2-HLIkByLd65vxjRD_kjE7gzYE3WAcGdxaQEEQNReY"
sheetName = "Data"
//barcode=6287029390129
function doGet(e) {
  barcode = e.parameter.barcode;
  // Open Google Sheet using ID
  var ss = SpreadsheetApp.openById(fileID);
  var sheet = ss.getSheetByName(sheetName);
  // Read all data rows from Google Sheet
  const values = sheet.getRange(2, 1, sheet.getLastRow() - 1, sheet.getLastColumn()).getValues();

const letterToColumn = letter => {
  let column = 0,
    length = letter.length;
  for (let i = 0; i < length; i++) {
    column += (letter.charCodeAt(i) - 64) * Math.pow(26, length - i - 1);
  }
  return column;
};

const columnLetters = ['W', 'X', 'Y', 'B', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'V', 'Z']; // Column letters you want to retrieve.
const fields = values.map(r => columnLetters.map(e => r[letterToColumn(e) - 1]));

const selected = fields.filter(line => line[0].toString() == barcode);

// Converts data rows in json format
const result = JSON.stringify({
        BarCode:              selected[0][0],
        SKUCode:              selected[0][1],
        VendorCode:           selected[0][2],
        RegistrationDate:     selected[0][3],
        VendorName:           selected[0][4],
        BrandName:            selected[0][5],
        ContactPerson:        selected[0][6],
        ContactNumber:        selected[0][7],
        ItemName:             selected[0][8],
        ItemImage:            selected[0][9],
        NetWeight:            selected[0][10],
        CartoonPack:          selected[0][11],
        StorageTemperature:   selected[0][12],
        ShelfLife:            selected[0][13],
        ShelfPrice:           selected[0][14],
        KottofCost:           selected[0][15],
        SupplyType:           selected[0][16],
        CoveredAreas:         selected[0][17],
        MinimumOrderQty:      selected[0][18],
        ContractDate:         selected[0][19],
        Notes:                selected[0][20],
  });

// Returns Result
return ContentService.createTextOutput(result).setMimeType(ContentService.MimeType.JSON);
}
