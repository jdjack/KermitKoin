package kermit;

import javafx.fxml.FXML;
import javafx.scene.input.Clipboard;
import javafx.scene.input.ClipboardContent;
import javafx.scene.text.Text;


public class ReceiveController {

  String address;
  @FXML
  private Text addressField;

  public void setAddress(String address) {
    this.address = address;
    addressField.setText(address);
  }

  @FXML
  private void handleCopyButtonAction() {
    final Clipboard clipboard = Clipboard.getSystemClipboard();
    final ClipboardContent content = new ClipboardContent();
    content.putString(address);
    clipboard.setContent(content);
  }
}
